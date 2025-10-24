// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateMmsJobResponseBodyData) *CreateMmsJobResponseBody
	GetData() *CreateMmsJobResponseBodyData
	SetRequestId(v string) *CreateMmsJobResponseBody
	GetRequestId() *string
}

type CreateMmsJobResponseBody struct {
	Data      *CreateMmsJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMmsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMmsJobResponseBody) GetData() *CreateMmsJobResponseBodyData {
	return s.Data
}

func (s *CreateMmsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMmsJobResponseBody) SetData(v *CreateMmsJobResponseBodyData) *CreateMmsJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateMmsJobResponseBody) SetRequestId(v string) *CreateMmsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMmsJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMmsJobResponseBodyData struct {
	AsyncTaskId *int64 `json:"asyncTaskId,omitempty" xml:"asyncTaskId,omitempty"`
}

func (s CreateMmsJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMmsJobResponseBodyData) GetAsyncTaskId() *int64 {
	return s.AsyncTaskId
}

func (s *CreateMmsJobResponseBodyData) SetAsyncTaskId(v int64) *CreateMmsJobResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *CreateMmsJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
