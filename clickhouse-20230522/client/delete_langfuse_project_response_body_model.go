// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLangfuseProjectResponseBody
	GetRequestId() *string
}

type DeleteLangfuseProjectResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLangfuseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLangfuseProjectResponseBody) SetRequestId(v string) *DeleteLangfuseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLangfuseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
