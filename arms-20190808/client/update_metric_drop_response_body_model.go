// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetricDropResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateMetricDropResponseBody
	GetCode() *int64
	SetData(v string) *UpdateMetricDropResponseBody
	GetData() *string
	SetMessage(v string) *UpdateMetricDropResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMetricDropResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMetricDropResponseBody
	GetSuccess() *bool
}

type UpdateMetricDropResponseBody struct {
	// The response status. Valid values: 2XX: The request is successful. 3XX: A redirection message is returned. 4XX: The request is invalid. 5XX: A server error occurs.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CCCA4B88-BD7B-5A38-89AF-C09293BD4187
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

func (s UpdateMetricDropResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetricDropResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetricDropResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateMetricDropResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateMetricDropResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMetricDropResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetricDropResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMetricDropResponseBody) SetCode(v int64) *UpdateMetricDropResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMetricDropResponseBody) SetData(v string) *UpdateMetricDropResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMetricDropResponseBody) SetMessage(v string) *UpdateMetricDropResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMetricDropResponseBody) SetRequestId(v string) *UpdateMetricDropResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetricDropResponseBody) SetSuccess(v bool) *UpdateMetricDropResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMetricDropResponseBody) Validate() error {
	return dara.Validate(s)
}
