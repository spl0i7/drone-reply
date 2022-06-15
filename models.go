package main

type MQTTFlightTelemetry struct {
	Bid  string `json:"bid"`
	Data struct {
		Five300 struct {
			GimbalPitch             float64 `json:"gimbal_pitch"`
			GimbalRoll              float64 `json:"gimbal_roll"`
			GimbalYaw               float64 `json:"gimbal_yaw"`
			MeasureTargetAltitude   float64 `json:"measure_target_altitude"`
			MeasureTargetDistance   float64 `json:"measure_target_distance"`
			MeasureTargetErrorState float64 `json:"measure_target_error_state"`
			MeasureTargetLatitude   float64 `json:"measure_target_latitude"`
			MeasureTargetLongitude  float64 `json:"measure_target_longitude"`
			PayloadIndex            string  `json:"payload_index"`
			Version                 int     `json:"version"`
		} `json:"53-0-0"`
		AttitudeHead  float64 `json:"attitude_head"`
		AttitudePitch float64 `json:"attitude_pitch"`
		AttitudeRoll  float64 `json:"attitude_roll"`
		Battery       struct {
			Batteries []struct {
				CapacityPercent int     `json:"capacity_percent"`
				FirmwareVersion string  `json:"firmware_version"`
				Index           int     `json:"index"`
				LoopTimes       int     `json:"loop_times"`
				Sn              string  `json:"sn"`
				SubType         int     `json:"sub_type"`
				Temperature     float64 `json:"temperature"`
				Type            int     `json:"type"`
				Voltage         int     `json:"voltage"`
			} `json:"batteries"`
			CapacityPercent  int `json:"capacity_percent"`
			LandingPower     int `json:"landing_power"`
			RemainFlightTime int `json:"remain_flight_time"`
			ReturnHomePower  int `json:"return_home_power"`
		} `json:"battery"`
		Elevation       float64 `json:"elevation"`
		FirmwareVersion string  `json:"firmware_version"`
		Gear            float64 `json:"gear"`
		Height          float64 `json:"height"`
		HomeDistance    float64 `json:"home_distance"`
		HorizontalSpeed float64 `json:"horizontal_speed"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
		ModeCode        float64 `json:"mode_code"`
		PositionState   struct {
			GpsNumber int `json:"gps_number"`
			IsFixed   int `json:"is_fixed"`
			Quality   int `json:"quality"`
			RtkNumber int `json:"rtk_number"`
		} `json:"position_state"`
		TotalFlightDistance float64 `json:"total_flight_distance"`
		TotalFlightTime     float64 `json:"total_flight_time"`
		VerticalSpeed       float64 `json:"vertical_speed"`
		WindDirection       float64 `json:"wind_direction"`
		WindSpeed           float64 `json:"wind_speed"`
	} `json:"data"`

	Tid       string `json:"tid"`
	Timestamp int64  `json:"timestamp"`
	Gateway   string `json:"gateway"`
}

type MQTTFlightStatus struct {
	Tid       string `json:"tid"`
	Bid       string `json:"bid"`
	Timestamp int64  `json:"timestamp"`
	Method    string `json:"method"`
	Data      struct {
		Type         int    `json:"type"`
		SubType      int    `json:"sub_type"`
		DeviceSecret string `json:"device_secret"`
		Nonce        string `json:"nonce"`
		Version      int    `json:"version"`
		SubDevices   []struct {
			Sn           string `json:"sn"`
			Type         int    `json:"type"`
			SubType      int    `json:"sub_type"`
			Index        string `json:"index"`
			DeviceSecret string `json:"device_secret"`
			Nonce        string `json:"nonce"`
			Version      int    `json:"version"`
		} `json:"sub_devices"`
	} `json:"data"`
}

type MQTTFlightStatusResponse struct {
	Tid  string `json:"tid"`
	Bid  string `json:"bid"`
	Data struct {
		Result int `json:"result"`
	} `json:"data"`
}
