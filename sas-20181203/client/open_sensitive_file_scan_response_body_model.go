// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSensitiveFileScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenSensitiveFileScanResponseBody
	GetCode() *string
	SetData(v *OpenSensitiveFileScanResponseBodyData) *OpenSensitiveFileScanResponseBody
	GetData() *OpenSensitiveFileScanResponseBodyData
	SetHttpStatusCode(v int32) *OpenSensitiveFileScanResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OpenSensitiveFileScanResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenSensitiveFileScanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OpenSensitiveFileScanResponseBody
	GetSuccess() *bool
}

type OpenSensitiveFileScanResponseBody struct {
	// The result code. A value of **200*	- indicates success. Any other value indicates failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned for modifying the sensitive file scan switch.
	Data *OpenSensitiveFileScanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The detailed information of the error code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B4A4C081-7F06-5481-9323-02A5419B9423
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result status of the API call. Valid values:
	//
	// - **true**: The API call was successful.
	//
	// - **false**: The API call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenSensitiveFileScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenSensitiveFileScanResponseBody) GoString() string {
	return s.String()
}

func (s *OpenSensitiveFileScanResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenSensitiveFileScanResponseBody) GetData() *OpenSensitiveFileScanResponseBodyData {
	return s.Data
}

func (s *OpenSensitiveFileScanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OpenSensitiveFileScanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenSensitiveFileScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenSensitiveFileScanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OpenSensitiveFileScanResponseBody) SetCode(v string) *OpenSensitiveFileScanResponseBody {
	s.Code = &v
	return s
}

func (s *OpenSensitiveFileScanResponseBody) SetData(v *OpenSensitiveFileScanResponseBodyData) *OpenSensitiveFileScanResponseBody {
	s.Data = v
	return s
}

func (s *OpenSensitiveFileScanResponseBody) SetHttpStatusCode(v int32) *OpenSensitiveFileScanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OpenSensitiveFileScanResponseBody) SetMessage(v string) *OpenSensitiveFileScanResponseBody {
	s.Message = &v
	return s
}

func (s *OpenSensitiveFileScanResponseBody) SetRequestId(v string) *OpenSensitiveFileScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenSensitiveFileScanResponseBody) SetSuccess(v bool) *OpenSensitiveFileScanResponseBody {
	s.Success = &v
	return s
}

func (s *OpenSensitiveFileScanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenSensitiveFileScanResponseBodyData struct {
	// The switch operation. Valid values:
	//
	// - **on**: Enable.
	//
	// - **off**: Disable.
	//
	// example:
	//
	// on
	SwitchOn *string `json:"SwitchOn,omitempty" xml:"SwitchOn,omitempty"`
}

func (s OpenSensitiveFileScanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OpenSensitiveFileScanResponseBodyData) GoString() string {
	return s.String()
}

func (s *OpenSensitiveFileScanResponseBodyData) GetSwitchOn() *string {
	return s.SwitchOn
}

func (s *OpenSensitiveFileScanResponseBodyData) SetSwitchOn(v string) *OpenSensitiveFileScanResponseBodyData {
	s.SwitchOn = &v
	return s
}

func (s *OpenSensitiveFileScanResponseBodyData) Validate() error {
	return dara.Validate(s)
}
