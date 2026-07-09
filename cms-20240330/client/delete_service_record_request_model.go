// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordType(v string) *DeleteServiceRecordRequest
	GetRecordType() *string
}

type DeleteServiceRecordRequest struct {
	// The type of the association entry. Valid values:
	//
	// logCorrelation: application log association
	//
	// This parameter is required.
	//
	// example:
	//
	// logCorrelation
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
}

func (s DeleteServiceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRecordRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *DeleteServiceRecordRequest) SetRecordType(v string) *DeleteServiceRecordRequest {
	s.RecordType = &v
	return s
}

func (s *DeleteServiceRecordRequest) Validate() error {
	return dara.Validate(s)
}
