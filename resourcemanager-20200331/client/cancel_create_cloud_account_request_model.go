// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCreateCloudAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *CancelCreateCloudAccountRequest
	GetRecordId() *string
}

type CancelCreateCloudAccountRequest struct {
	// The account record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06950264-3f0d-4ca9-82dd-6ee7a3d33d6b
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s CancelCreateCloudAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCreateCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *CancelCreateCloudAccountRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *CancelCreateCloudAccountRequest) SetRecordId(v string) *CancelCreateCloudAccountRequest {
	s.RecordId = &v
	return s
}

func (s *CancelCreateCloudAccountRequest) Validate() error {
	return dara.Validate(s)
}
