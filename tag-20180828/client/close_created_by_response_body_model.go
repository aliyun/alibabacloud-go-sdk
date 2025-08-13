// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseCreatedByResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseCreatedByResponseBody
	GetRequestId() *string
}

type CloseCreatedByResponseBody struct {
	// example:
	//
	// AECFE0F2-CEC3-5D16-BE4C-E2F95083D063
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseCreatedByResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseCreatedByResponseBody) GoString() string {
	return s.String()
}

func (s *CloseCreatedByResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseCreatedByResponseBody) SetRequestId(v string) *CloseCreatedByResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseCreatedByResponseBody) Validate() error {
	return dara.Validate(s)
}
