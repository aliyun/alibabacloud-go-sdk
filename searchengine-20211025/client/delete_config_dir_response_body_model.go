// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigDirResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConfigDirResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DeleteConfigDirResponseBody
	GetResult() map[string]interface{}
}

type DeleteConfigDirResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F43E8AB4-419C-5F4C-90D6-615590DFAA3C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteConfigDirResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigDirResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigDirResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConfigDirResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DeleteConfigDirResponseBody) SetRequestId(v string) *DeleteConfigDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConfigDirResponseBody) SetResult(v map[string]interface{}) *DeleteConfigDirResponseBody {
	s.Result = v
	return s
}

func (s *DeleteConfigDirResponseBody) Validate() error {
	return dara.Validate(s)
}
