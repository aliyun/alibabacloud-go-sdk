// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotPresetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHoneypotPresetResponseBody
	GetCode() *string
	SetHoneypotPreset(v *CreateHoneypotPresetResponseBodyHoneypotPreset) *CreateHoneypotPresetResponseBody
	GetHoneypotPreset() *CreateHoneypotPresetResponseBodyHoneypotPreset
	SetHttpStatusCode(v int32) *CreateHoneypotPresetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateHoneypotPresetResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHoneypotPresetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHoneypotPresetResponseBody
	GetSuccess() *bool
}

type CreateHoneypotPresetResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation result.
	HoneypotPreset *CreateHoneypotPresetResponseBodyHoneypotPreset `json:"HoneypotPreset,omitempty" xml:"HoneypotPreset,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7C1C6E65-C8B9-54C9-9F92-2F5E51C4E16D
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

func (s CreateHoneypotPresetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotPresetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHoneypotPresetResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHoneypotPresetResponseBody) GetHoneypotPreset() *CreateHoneypotPresetResponseBodyHoneypotPreset {
	return s.HoneypotPreset
}

func (s *CreateHoneypotPresetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateHoneypotPresetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHoneypotPresetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHoneypotPresetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHoneypotPresetResponseBody) SetCode(v string) *CreateHoneypotPresetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHoneypotPresetResponseBody) SetHoneypotPreset(v *CreateHoneypotPresetResponseBodyHoneypotPreset) *CreateHoneypotPresetResponseBody {
	s.HoneypotPreset = v
	return s
}

func (s *CreateHoneypotPresetResponseBody) SetHttpStatusCode(v int32) *CreateHoneypotPresetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateHoneypotPresetResponseBody) SetMessage(v string) *CreateHoneypotPresetResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHoneypotPresetResponseBody) SetRequestId(v string) *CreateHoneypotPresetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHoneypotPresetResponseBody) SetSuccess(v bool) *CreateHoneypotPresetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHoneypotPresetResponseBody) Validate() error {
	if s.HoneypotPreset != nil {
		if err := s.HoneypotPreset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateHoneypotPresetResponseBodyHoneypotPreset struct {
	// The ID of the honeypot template.
	//
	// example:
	//
	// d6ece172-34d9-4942-99a4-b309cb55xxxx
	HoneypotPresetId *string `json:"HoneypotPresetId,omitempty" xml:"HoneypotPresetId,omitempty"`
}

func (s CreateHoneypotPresetResponseBodyHoneypotPreset) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotPresetResponseBodyHoneypotPreset) GoString() string {
	return s.String()
}

func (s *CreateHoneypotPresetResponseBodyHoneypotPreset) GetHoneypotPresetId() *string {
	return s.HoneypotPresetId
}

func (s *CreateHoneypotPresetResponseBodyHoneypotPreset) SetHoneypotPresetId(v string) *CreateHoneypotPresetResponseBodyHoneypotPreset {
	s.HoneypotPresetId = &v
	return s
}

func (s *CreateHoneypotPresetResponseBodyHoneypotPreset) Validate() error {
	return dara.Validate(s)
}
