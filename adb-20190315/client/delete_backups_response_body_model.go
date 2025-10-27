// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBackupsResponseBody
	GetRequestId() *string
}

type DeleteBackupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 771C5FAA-530F-52F7-B84D-EBAD4561D590
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupsResponseBody) SetRequestId(v string) *DeleteBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupsResponseBody) Validate() error {
	return dara.Validate(s)
}
