// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperateCommonOverallConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchOperateCommonOverallConfigResponseBody
	GetRequestId() *string
}

type BatchOperateCommonOverallConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 765EDBDE-1686-5DBA-B76F-2E0E6E7E1B96
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchOperateCommonOverallConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchOperateCommonOverallConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchOperateCommonOverallConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchOperateCommonOverallConfigResponseBody) SetRequestId(v string) *BatchOperateCommonOverallConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchOperateCommonOverallConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
