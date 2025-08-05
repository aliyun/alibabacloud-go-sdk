// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateHanaInstanceResponseBody
	GetClusterId() *string
	SetCode(v string) *CreateHanaInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *CreateHanaInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHanaInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHanaInstanceResponseBody
	GetSuccess() *bool
}

type CreateHanaInstanceResponseBody struct {
	// The ID of the SAP HANA instance.
	//
	// example:
	//
	// cl-000dp1sz******6hn
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EB526A5D-1FE2-51C1-B790-1732C1DBA969
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHanaInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHanaInstanceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateHanaInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHanaInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHanaInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHanaInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHanaInstanceResponseBody) SetClusterId(v string) *CreateHanaInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) SetCode(v string) *CreateHanaInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) SetMessage(v string) *CreateHanaInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) SetRequestId(v string) *CreateHanaInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) SetSuccess(v bool) *CreateHanaInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHanaInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
