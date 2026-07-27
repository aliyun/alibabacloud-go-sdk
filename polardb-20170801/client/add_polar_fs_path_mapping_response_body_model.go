// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarFsPathMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddPolarFsPathMappingResponseBody
	GetRequestId() *string
}

type AddPolarFsPathMappingResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPolarFsPathMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPolarFsPathMappingResponseBody) GoString() string {
	return s.String()
}

func (s *AddPolarFsPathMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPolarFsPathMappingResponseBody) SetRequestId(v string) *AddPolarFsPathMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPolarFsPathMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
