// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileProtectRemarkResponseBody
	GetRequestId() *string
}

type UpdateFileProtectRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 226440DF-DFCD-5B93-9951-FCF0A16A6B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileProtectRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileProtectRemarkResponseBody) SetRequestId(v string) *UpdateFileProtectRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileProtectRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
