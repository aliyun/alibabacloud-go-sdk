// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBizMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBizMetricResponseBody
	GetCode() *string
	SetData(v *CreateBizMetricResponseBodyData) *CreateBizMetricResponseBody
	GetData() *CreateBizMetricResponseBodyData
	SetHttpStatusCode(v int32) *CreateBizMetricResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBizMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBizMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBizMetricResponseBody
	GetSuccess() *bool
}

type CreateBizMetricResponseBody struct {
	// example:
	//
	// success
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateBizMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBizMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBizMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBizMetricResponseBody) GetData() *CreateBizMetricResponseBodyData {
	return s.Data
}

func (s *CreateBizMetricResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBizMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBizMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBizMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBizMetricResponseBody) SetCode(v string) *CreateBizMetricResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBizMetricResponseBody) SetData(v *CreateBizMetricResponseBodyData) *CreateBizMetricResponseBody {
	s.Data = v
	return s
}

func (s *CreateBizMetricResponseBody) SetHttpStatusCode(v int32) *CreateBizMetricResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBizMetricResponseBody) SetMessage(v string) *CreateBizMetricResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBizMetricResponseBody) SetRequestId(v string) *CreateBizMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBizMetricResponseBody) SetSuccess(v bool) *CreateBizMetricResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBizMetricResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBizMetricResponseBodyData struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBizMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateBizMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBizMetricResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateBizMetricResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBizMetricResponseBodyData) SetMessage(v string) *CreateBizMetricResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateBizMetricResponseBodyData) SetSuccess(v bool) *CreateBizMetricResponseBodyData {
	s.Success = &v
	return s
}

func (s *CreateBizMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
