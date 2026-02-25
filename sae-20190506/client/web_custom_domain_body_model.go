// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebCustomDomainBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebCustomDomainBody
	GetCode() *int32
	SetData(v *WebCustomDomain) *WebCustomDomainBody
	GetData() *WebCustomDomain
	SetMessage(v string) *WebCustomDomainBody
	GetMessage() *string
	SetRequestId(v string) *WebCustomDomainBody
	GetRequestId() *string
	SetSuccess(v bool) *WebCustomDomainBody
	GetSuccess() *bool
}

type WebCustomDomainBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned for the operation.
	Data *WebCustomDomain `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned for the operation. Valid values:
	//
	// 	- success is returned when the request succeeds.
	//
	// 	- An error code is returned if the request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WebCustomDomainBody) String() string {
	return dara.Prettify(s)
}

func (s WebCustomDomainBody) GoString() string {
	return s.String()
}

func (s *WebCustomDomainBody) GetCode() *int32 {
	return s.Code
}

func (s *WebCustomDomainBody) GetData() *WebCustomDomain {
	return s.Data
}

func (s *WebCustomDomainBody) GetMessage() *string {
	return s.Message
}

func (s *WebCustomDomainBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebCustomDomainBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebCustomDomainBody) SetCode(v int32) *WebCustomDomainBody {
	s.Code = &v
	return s
}

func (s *WebCustomDomainBody) SetData(v *WebCustomDomain) *WebCustomDomainBody {
	s.Data = v
	return s
}

func (s *WebCustomDomainBody) SetMessage(v string) *WebCustomDomainBody {
	s.Message = &v
	return s
}

func (s *WebCustomDomainBody) SetRequestId(v string) *WebCustomDomainBody {
	s.RequestId = &v
	return s
}

func (s *WebCustomDomainBody) SetSuccess(v bool) *WebCustomDomainBody {
	s.Success = &v
	return s
}

func (s *WebCustomDomainBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
