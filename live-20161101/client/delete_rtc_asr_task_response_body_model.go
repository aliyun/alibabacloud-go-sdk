// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtcAsrTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DeleteRtcAsrTaskResponseBody
	GetDescription() *string
	SetRequestId(v string) *DeleteRtcAsrTaskResponseBody
	GetRequestId() *string
	SetRetCode(v int64) *DeleteRtcAsrTaskResponseBody
	GetRetCode() *int64
}

type DeleteRtcAsrTaskResponseBody struct {
	// The result of the request. If success is returned, the request is successful. If an error message is returned, the request failed.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 85F94125-B695-1FB8-A7E7-3BE7CE07EF31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned HTTP status code. HTTP status code 2000 indicates that the request is successful. If another HTTP status code is returned, the request failed.
	//
	// example:
	//
	// 2000
	RetCode *int64 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s DeleteRtcAsrTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtcAsrTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRtcAsrTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DeleteRtcAsrTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRtcAsrTaskResponseBody) GetRetCode() *int64 {
	return s.RetCode
}

func (s *DeleteRtcAsrTaskResponseBody) SetDescription(v string) *DeleteRtcAsrTaskResponseBody {
	s.Description = &v
	return s
}

func (s *DeleteRtcAsrTaskResponseBody) SetRequestId(v string) *DeleteRtcAsrTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRtcAsrTaskResponseBody) SetRetCode(v int64) *DeleteRtcAsrTaskResponseBody {
	s.RetCode = &v
	return s
}

func (s *DeleteRtcAsrTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
