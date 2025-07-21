// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagResourceResponseBody
	GetRequestId() *string
}

type TagResourceResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4162a6af-bc99-40b3-a552-89dcc8aaf7c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourceResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourceResponseBody) SetRequestId(v string) *TagResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
