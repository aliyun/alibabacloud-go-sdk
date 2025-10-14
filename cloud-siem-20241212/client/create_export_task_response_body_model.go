// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateExportTaskResponseBody
	GetFileName() *string
	SetId(v int64) *CreateExportTaskResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateExportTaskResponseBody
	GetRequestId() *string
}

type CreateExportTaskResponseBody struct {
	// example:
	//
	// event_1193842352994186_17344888****.xlsx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 400151
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExportTaskResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *CreateExportTaskResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExportTaskResponseBody) SetFileName(v string) *CreateExportTaskResponseBody {
	s.FileName = &v
	return s
}

func (s *CreateExportTaskResponseBody) SetId(v int64) *CreateExportTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateExportTaskResponseBody) SetRequestId(v string) *CreateExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
