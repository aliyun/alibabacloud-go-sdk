// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImportLogTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitImportLogTasksResponseBodyData) *SubmitImportLogTasksResponseBody
	GetData() *SubmitImportLogTasksResponseBodyData
	SetRequestId(v string) *SubmitImportLogTasksResponseBody
	GetRequestId() *string
}

type SubmitImportLogTasksResponseBody struct {
	// The data returned.
	Data *SubmitImportLogTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitImportLogTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitImportLogTasksResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImportLogTasksResponseBody) GetData() *SubmitImportLogTasksResponseBodyData {
	return s.Data
}

func (s *SubmitImportLogTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitImportLogTasksResponseBody) SetData(v *SubmitImportLogTasksResponseBodyData) *SubmitImportLogTasksResponseBody {
	s.Data = v
	return s
}

func (s *SubmitImportLogTasksResponseBody) SetRequestId(v string) *SubmitImportLogTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImportLogTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitImportLogTasksResponseBodyData struct {
	// The number of log collection tasks that are submitted.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s SubmitImportLogTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitImportLogTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitImportLogTasksResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *SubmitImportLogTasksResponseBodyData) SetCount(v int32) *SubmitImportLogTasksResponseBodyData {
	s.Count = &v
	return s
}

func (s *SubmitImportLogTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
