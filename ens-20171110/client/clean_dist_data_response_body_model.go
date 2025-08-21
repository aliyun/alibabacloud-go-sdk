// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanDistDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CleanDistDataResponseBody
	GetRequestId() *string
}

type CleanDistDataResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CleanDistDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CleanDistDataResponseBody) GoString() string {
	return s.String()
}

func (s *CleanDistDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CleanDistDataResponseBody) SetRequestId(v string) *CleanDistDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CleanDistDataResponseBody) Validate() error {
	return dara.Validate(s)
}
