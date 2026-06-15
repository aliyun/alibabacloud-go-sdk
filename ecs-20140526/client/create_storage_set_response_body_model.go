// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateStorageSetResponseBody
	GetRequestId() *string
	SetStorageSetId(v string) *CreateStorageSetResponseBody
	GetStorageSetId() *string
}

type CreateStorageSetResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
}

func (s CreateStorageSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStorageSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStorageSetResponseBody) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *CreateStorageSetResponseBody) SetRequestId(v string) *CreateStorageSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStorageSetResponseBody) SetStorageSetId(v string) *CreateStorageSetResponseBody {
	s.StorageSetId = &v
	return s
}

func (s *CreateStorageSetResponseBody) Validate() error {
	return dara.Validate(s)
}
