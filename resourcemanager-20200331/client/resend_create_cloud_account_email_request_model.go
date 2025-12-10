// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendCreateCloudAccountEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *ResendCreateCloudAccountEmailRequest
	GetRecordId() *string
}

type ResendCreateCloudAccountEmailRequest struct {
	// The account record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06950264-3f0d-4ca9-82dd-6ee7a3d33d6b
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s ResendCreateCloudAccountEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ResendCreateCloudAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *ResendCreateCloudAccountEmailRequest) SetRecordId(v string) *ResendCreateCloudAccountEmailRequest {
	s.RecordId = &v
	return s
}

func (s *ResendCreateCloudAccountEmailRequest) Validate() error {
	return dara.Validate(s)
}
