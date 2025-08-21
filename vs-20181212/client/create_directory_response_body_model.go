// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateDirectoryResponseBody
	GetId() *string
	SetRequestId(v string) *CreateDirectoryResponseBody
	GetRequestId() *string
}

type CreateDirectoryResponseBody struct {
	// example:
	//
	// 399*****488-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDirectoryResponseBody) SetId(v string) *CreateDirectoryResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDirectoryResponseBody) SetRequestId(v string) *CreateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
