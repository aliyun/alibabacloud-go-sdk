// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendPromoteResourceAccountEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *ResendPromoteResourceAccountEmailRequest
	GetRecordId() *string
}

type ResendPromoteResourceAccountEmailRequest struct {
	// The account record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06950264-3f0d-4ca9-82dd-6ee7a3d33d6b
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s ResendPromoteResourceAccountEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailRequest) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *ResendPromoteResourceAccountEmailRequest) SetRecordId(v string) *ResendPromoteResourceAccountEmailRequest {
	s.RecordId = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailRequest) Validate() error {
	return dara.Validate(s)
}
