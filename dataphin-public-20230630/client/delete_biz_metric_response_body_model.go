// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBizMetricResponseBody
	GetCode() *string
	SetData(v *DeleteBizMetricResponseBodyData) *DeleteBizMetricResponseBody
	GetData() *DeleteBizMetricResponseBodyData
	SetHttpStatusCode(v int32) *DeleteBizMetricResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteBizMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBizMetricResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBizMetricResponseBody
	GetSuccess() *bool
}

type DeleteBizMetricResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteBizMetricResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBizMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBizMetricResponseBody) GetData() *DeleteBizMetricResponseBodyData {
	return s.Data
}

func (s *DeleteBizMetricResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBizMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBizMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBizMetricResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBizMetricResponseBody) SetCode(v string) *DeleteBizMetricResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBizMetricResponseBody) SetData(v *DeleteBizMetricResponseBodyData) *DeleteBizMetricResponseBody {
	s.Data = v
	return s
}

func (s *DeleteBizMetricResponseBody) SetHttpStatusCode(v int32) *DeleteBizMetricResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBizMetricResponseBody) SetMessage(v string) *DeleteBizMetricResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBizMetricResponseBody) SetRequestId(v string) *DeleteBizMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBizMetricResponseBody) SetSuccess(v bool) *DeleteBizMetricResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBizMetricResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteBizMetricResponseBodyData struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBizMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteBizMetricResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteBizMetricResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBizMetricResponseBodyData) SetMessage(v string) *DeleteBizMetricResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteBizMetricResponseBodyData) SetSuccess(v bool) *DeleteBizMetricResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteBizMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}
