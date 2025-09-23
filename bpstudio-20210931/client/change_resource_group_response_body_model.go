// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeResourceGroupResponseBody
	GetCode() *string
	SetData(v string) *ChangeResourceGroupResponseBody
	GetData() *string
	SetMessage(v string) *ChangeResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
}

type ChangeResourceGroupResponseBody struct {
	// The HTTP status code. A value of 200 indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// No business data is returned for this parameter.
	//
	// example:
	//
	// No business data is returned for this parameter.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3488F1D-C444-17D0-BA4F-5374BA0F3562
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeResourceGroupResponseBody) GetData() *string {
	return s.Data
}

func (s *ChangeResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) SetCode(v string) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetData(v string) *ChangeResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
