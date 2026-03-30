// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreAccessKeyFromRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestoreAccessKeyFromRecycleBinResponseBody
	GetRequestId() *string
}

type RestoreAccessKeyFromRecycleBinResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreAccessKeyFromRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreAccessKeyFromRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreAccessKeyFromRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreAccessKeyFromRecycleBinResponseBody) SetRequestId(v string) *RestoreAccessKeyFromRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinResponseBody) Validate() error {
	return dara.Validate(s)
}
