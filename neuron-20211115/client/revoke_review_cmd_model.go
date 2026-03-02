// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeReviewCmd interface {
	dara.Model
	String() string
	GoString() string
	SetPbcVersionId(v int64) *RevokeReviewCmd
	GetPbcVersionId() *int64
}

type RevokeReviewCmd struct {
	PbcVersionId *int64 `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
}

func (s RevokeReviewCmd) String() string {
	return dara.Prettify(s)
}

func (s RevokeReviewCmd) GoString() string {
	return s.String()
}

func (s *RevokeReviewCmd) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *RevokeReviewCmd) SetPbcVersionId(v int64) *RevokeReviewCmd {
	s.PbcVersionId = &v
	return s
}

func (s *RevokeReviewCmd) Validate() error {
	return dara.Validate(s)
}
