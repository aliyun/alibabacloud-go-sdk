// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPromoteResourceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *CancelPromoteResourceAccountRequest
	GetRecordId() *string
}

type CancelPromoteResourceAccountRequest struct {
	// The account record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06950264-3f0d-4ca9-82dd-6ee7a3d33d6b
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s CancelPromoteResourceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPromoteResourceAccountRequest) GoString() string {
	return s.String()
}

func (s *CancelPromoteResourceAccountRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *CancelPromoteResourceAccountRequest) SetRecordId(v string) *CancelPromoteResourceAccountRequest {
	s.RecordId = &v
	return s
}

func (s *CancelPromoteResourceAccountRequest) Validate() error {
	return dara.Validate(s)
}
