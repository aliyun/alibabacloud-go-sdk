// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTokenResponseBody
	GetRequestId() *string
	SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody
	GetData() *GetTokenResponseBodyData
	SetErrorCode(v string) *GetTokenResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *GetTokenResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *GetTokenResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *GetTokenResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *GetTokenResponseBody
	GetSuccess() *bool
}

type GetTokenResponseBody struct {
	// Request RequestId
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Correctly processed return data
	Data *GetTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Business error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// Data carried during error handling
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// Error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// When the HTTP request is successful, the status value is 200.
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// Whether it is correct
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenResponseBody) GetData() *GetTokenResponseBodyData {
	return s.Data
}

func (s *GetTokenResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTokenResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *GetTokenResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetTokenResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetErrorCode(v string) *GetTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTokenResponseBody) SetErrorData(v interface{}) *GetTokenResponseBody {
	s.ErrorData = v
	return s
}

func (s *GetTokenResponseBody) SetErrorMsg(v string) *GetTokenResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTokenResponseBody) SetStatus(v int32) *GetTokenResponseBody {
	s.Status = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v bool) *GetTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTokenResponseBodyData struct {
	// Remaining valid time of the token in seconds
	//
	// example:
	//
	// 7200
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// Timestamp of token generation in seconds
	//
	// example:
	//
	// 1677055176
	GenerateTime *int64 `json:"generate_time,omitempty" xml:"generate_time,omitempty"`
	// token
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiIyUlRjY0Ezc1puSndpYU11R1ctTkVRIiwiaWF0IjoxNjc3MTU1Njg3LCJleHAiOjE2NzcxNjI4ODcsIm5iZiI6MTY3NzE1NTYyN30.bd8qBedJ7R24NC8VpMtM4Ni5OR-Cc0utPiKSx8fjoj9taalt7zXBF8uIVTETw1N-Fx9reLflwVXrbDHky7ZKqlE5o_B5Bkx-crQKlJL-NzasEmNnuJNb5kmmbCy3mvIrQfo5Q5Y0ZaQ110pXK4qd1shRbklvuQXn8lPueJtNqo8VdIOAPGG_rPwOw2P767I0fyFHcome8FR4ST1UrwxeApNKMB_BkpCsUZLgpm9h9trhKbB-3qtk6UK1GKmfw6qlWpL3PQN7FAObNruS0r0CGh3Muc9PaGsuu8Xu5on21h9WmI7L0-jatZkS55p4PEerU56XpkwJfa35_hltKNTauu
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetTokenResponseBodyData) GetGenerateTime() *int64 {
	return s.GenerateTime
}

func (s *GetTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetTokenResponseBodyData) SetExpireTime(v int64) *GetTokenResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetGenerateTime(v int64) *GetTokenResponseBodyData {
	s.GenerateTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetToken(v string) *GetTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
