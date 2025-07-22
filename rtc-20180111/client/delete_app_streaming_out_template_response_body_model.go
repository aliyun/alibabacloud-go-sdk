// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppStreamingOutTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppStreamingOutTemplateResponseBody
	GetRequestId() *string
}

type DeleteAppStreamingOutTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD825C17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppStreamingOutTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppStreamingOutTemplateResponseBody) SetRequestId(v string) *DeleteAppStreamingOutTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
