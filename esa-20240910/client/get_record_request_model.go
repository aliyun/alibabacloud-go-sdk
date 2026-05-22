// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *GetRecordRequest
	GetRecordId() *int64
}

type GetRecordRequest struct {
	// This parameter is required.
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s GetRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecordRequest) GoString() string {
	return s.String()
}

func (s *GetRecordRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *GetRecordRequest) SetRecordId(v int64) *GetRecordRequest {
	s.RecordId = &v
	return s
}

func (s *GetRecordRequest) Validate() error {
	return dara.Validate(s)
}
