// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHoneypotProbeResponseBody
	GetCode() *string
	SetHoneypotProbe(v *CreateHoneypotProbeResponseBodyHoneypotProbe) *CreateHoneypotProbeResponseBody
	GetHoneypotProbe() *CreateHoneypotProbeResponseBodyHoneypotProbe
	SetHttpStatusCode(v int32) *CreateHoneypotProbeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateHoneypotProbeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHoneypotProbeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHoneypotProbeResponseBody
	GetSuccess() *bool
}

type CreateHoneypotProbeResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the probe.
	HoneypotProbe *CreateHoneypotProbeResponseBodyHoneypotProbe `json:"HoneypotProbe,omitempty" xml:"HoneypotProbe,omitempty" type:"Struct"`
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
	// B7A839E8-70AE-591D-8D9E-C5419A2240DB
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

func (s CreateHoneypotProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHoneypotProbeResponseBody) GetHoneypotProbe() *CreateHoneypotProbeResponseBodyHoneypotProbe {
	return s.HoneypotProbe
}

func (s *CreateHoneypotProbeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateHoneypotProbeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHoneypotProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHoneypotProbeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHoneypotProbeResponseBody) SetCode(v string) *CreateHoneypotProbeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHoneypotProbeResponseBody) SetHoneypotProbe(v *CreateHoneypotProbeResponseBodyHoneypotProbe) *CreateHoneypotProbeResponseBody {
	s.HoneypotProbe = v
	return s
}

func (s *CreateHoneypotProbeResponseBody) SetHttpStatusCode(v int32) *CreateHoneypotProbeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateHoneypotProbeResponseBody) SetMessage(v string) *CreateHoneypotProbeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHoneypotProbeResponseBody) SetRequestId(v string) *CreateHoneypotProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHoneypotProbeResponseBody) SetSuccess(v bool) *CreateHoneypotProbeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHoneypotProbeResponseBody) Validate() error {
	if s.HoneypotProbe != nil {
		if err := s.HoneypotProbe.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateHoneypotProbeResponseBodyHoneypotProbe struct {
	// The ID of the probe.
	//
	// example:
	//
	// b69e9aa8-2ea8-4c5a-836a-c1fbacff****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
}

func (s CreateHoneypotProbeResponseBodyHoneypotProbe) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeResponseBodyHoneypotProbe) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeResponseBodyHoneypotProbe) GetProbeId() *string {
	return s.ProbeId
}

func (s *CreateHoneypotProbeResponseBodyHoneypotProbe) SetProbeId(v string) *CreateHoneypotProbeResponseBodyHoneypotProbe {
	s.ProbeId = &v
	return s
}

func (s *CreateHoneypotProbeResponseBodyHoneypotProbe) Validate() error {
	return dara.Validate(s)
}
