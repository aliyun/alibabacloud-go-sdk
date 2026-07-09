// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordType(v string) *GetServiceRecordRequest
	GetRecordType() *string
}

type GetServiceRecordRequest struct {
	// The type of the linked entry. Currently supported values:
	//
	// logCorrelation: indicates application log association.
	//
	// This parameter is required.
	//
	// example:
	//
	// logCorrelation
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
}

func (s GetServiceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRecordRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRecordRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *GetServiceRecordRequest) SetRecordType(v string) *GetServiceRecordRequest {
	s.RecordType = &v
	return s
}

func (s *GetServiceRecordRequest) Validate() error {
	return dara.Validate(s)
}
