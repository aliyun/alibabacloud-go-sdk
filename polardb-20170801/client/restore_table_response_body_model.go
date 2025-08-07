// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestoreTableResponseBody
	GetRequestId() *string
}

type RestoreTableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C47508C-9DC8-455B-985E-2F2FA8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreTableResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreTableResponseBody) SetRequestId(v string) *RestoreTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreTableResponseBody) Validate() error {
	return dara.Validate(s)
}
