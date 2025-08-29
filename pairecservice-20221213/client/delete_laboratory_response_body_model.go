// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaboratoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLaboratoryResponseBody
	GetRequestId() *string
}

type DeleteLaboratoryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1C0898E5-9220-5443-B2D9-445FF0688215
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLaboratoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLaboratoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLaboratoryResponseBody) SetRequestId(v string) *DeleteLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLaboratoryResponseBody) Validate() error {
	return dara.Validate(s)
}
