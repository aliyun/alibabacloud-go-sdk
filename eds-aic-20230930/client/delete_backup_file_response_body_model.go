// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBackupFileResponseBody
	GetRequestId() *string
}

type DeleteBackupFileResponseBody struct {
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupFileResponseBody) SetRequestId(v string) *DeleteBackupFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupFileResponseBody) Validate() error {
	return dara.Validate(s)
}
