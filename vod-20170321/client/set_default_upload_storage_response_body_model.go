// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultUploadStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultUploadStorageResponseBody
	GetRequestId() *string
}

type SetDefaultUploadStorageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultUploadStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultUploadStorageResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultUploadStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultUploadStorageResponseBody) SetRequestId(v string) *SetDefaultUploadStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultUploadStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
