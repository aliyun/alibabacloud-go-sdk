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
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *OpenSensitiveFileScanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// B4A4C081-7F06-5481-9323-02A5419B9423
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
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
	// Indicates whether sensitive file scan is enabled or disabled. Valid values:
	//
	// 	- **on**: enabled
	//
	// 	- **off**: disabled
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
