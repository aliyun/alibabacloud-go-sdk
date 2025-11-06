// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppViewTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppViewTemplateResponseBody
	GetRequestId() *string
}

type DeleteAppViewTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD825C17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppViewTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppViewTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppViewTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppViewTemplateResponseBody) SetRequestId(v string) *DeleteAppViewTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppViewTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
