// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateMmsDataSourceResponseBodyData) *UpdateMmsDataSourceResponseBody
	GetData() *UpdateMmsDataSourceResponseBodyData
	SetRequestId(v string) *UpdateMmsDataSourceResponseBody
	GetRequestId() *string
}

type UpdateMmsDataSourceResponseBody struct {
	Data *UpdateMmsDataSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 76CE80C8-7392-5591-BCC8-610AFBF78ADF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMmsDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmsDataSourceResponseBody) GetData() *UpdateMmsDataSourceResponseBodyData {
	return s.Data
}

func (s *UpdateMmsDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmsDataSourceResponseBody) SetData(v *UpdateMmsDataSourceResponseBodyData) *UpdateMmsDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMmsDataSourceResponseBody) SetRequestId(v string) *UpdateMmsDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmsDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateMmsDataSourceResponseBodyData struct {
	AsyncTaskId *int64 `json:"asyncTaskId,omitempty" xml:"asyncTaskId,omitempty"`
	SourceId    *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s UpdateMmsDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMmsDataSourceResponseBodyData) GetAsyncTaskId() *int64 {
	return s.AsyncTaskId
}

func (s *UpdateMmsDataSourceResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *UpdateMmsDataSourceResponseBodyData) SetAsyncTaskId(v int64) *UpdateMmsDataSourceResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *UpdateMmsDataSourceResponseBodyData) SetSourceId(v int64) *UpdateMmsDataSourceResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *UpdateMmsDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
