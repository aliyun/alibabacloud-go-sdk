// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRumAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRumAppResponseBody
	GetCode() *string
	SetData(v *UpdateRumAppResponseBodyData) *UpdateRumAppResponseBody
	GetData() *UpdateRumAppResponseBodyData
	SetHttpStatusCode(v string) *UpdateRumAppResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *UpdateRumAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRumAppResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateRumAppResponseBody
	GetSuccess() *string
}

type UpdateRumAppResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the returned results.
	Data *UpdateRumAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9C9DA3D-10FE-472E-9EEF-2D0A3E41****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRumAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRumAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRumAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRumAppResponseBody) GetData() *UpdateRumAppResponseBodyData {
	return s.Data
}

func (s *UpdateRumAppResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateRumAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRumAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRumAppResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateRumAppResponseBody) SetCode(v string) *UpdateRumAppResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRumAppResponseBody) SetData(v *UpdateRumAppResponseBodyData) *UpdateRumAppResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRumAppResponseBody) SetHttpStatusCode(v string) *UpdateRumAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRumAppResponseBody) SetMessage(v string) *UpdateRumAppResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRumAppResponseBody) SetRequestId(v string) *UpdateRumAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRumAppResponseBody) SetSuccess(v string) *UpdateRumAppResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRumAppResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRumAppResponseBodyData struct {
	// The user configurations. This is a reserved parameter.
	//
	// example:
	//
	// null
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The QPS limit. Unit: bytes.
	//
	// example:
	//
	// 100000
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// Indicates whether the request is throttled due to the QPS limit. Valid values: true and false.
	//
	// example:
	//
	// true
	Limited *bool `json:"Limited,omitempty" xml:"Limited,omitempty"`
	// The usage. Unit: bytes.
	//
	// example:
	//
	// 10000
	Usage *int32 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s UpdateRumAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateRumAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRumAppResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *UpdateRumAppResponseBodyData) GetLimit() *int32 {
	return s.Limit
}

func (s *UpdateRumAppResponseBodyData) GetLimited() *bool {
	return s.Limited
}

func (s *UpdateRumAppResponseBodyData) GetUsage() *int32 {
	return s.Usage
}

func (s *UpdateRumAppResponseBodyData) SetConfig(v string) *UpdateRumAppResponseBodyData {
	s.Config = &v
	return s
}

func (s *UpdateRumAppResponseBodyData) SetLimit(v int32) *UpdateRumAppResponseBodyData {
	s.Limit = &v
	return s
}

func (s *UpdateRumAppResponseBodyData) SetLimited(v bool) *UpdateRumAppResponseBodyData {
	s.Limited = &v
	return s
}

func (s *UpdateRumAppResponseBodyData) SetUsage(v int32) *UpdateRumAppResponseBodyData {
	s.Usage = &v
	return s
}

func (s *UpdateRumAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
