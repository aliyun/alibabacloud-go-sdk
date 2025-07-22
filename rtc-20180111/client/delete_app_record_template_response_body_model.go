// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppRecordTemplateResponseBody
	GetRequestId() *string
}

type DeleteAppRecordTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppRecordTemplateResponseBody) SetRequestId(v string) *DeleteAppRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
