// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsPathMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePolarFsPathMappingResponseBody
	GetRequestId() *string
}

type DeletePolarFsPathMappingResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7E2FE3BB-C677-5FF9-9FC5-9CF364BD6BE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolarFsPathMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsPathMappingResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolarFsPathMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolarFsPathMappingResponseBody) SetRequestId(v string) *DeletePolarFsPathMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolarFsPathMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
