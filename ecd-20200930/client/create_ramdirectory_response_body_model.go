// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRAMDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *CreateRAMDirectoryResponseBody
	GetDirectoryId() *string
	SetRequestId(v string) *CreateRAMDirectoryResponseBody
	GetRequestId() *string
}

type CreateRAMDirectoryResponseBody struct {
	// The RAM directory ID.
	//
	// example:
	//
	// dri-uf62w3qzt4aigvlcb****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRAMDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRAMDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateRAMDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRAMDirectoryResponseBody) SetDirectoryId(v string) *CreateRAMDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateRAMDirectoryResponseBody) SetRequestId(v string) *CreateRAMDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRAMDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
