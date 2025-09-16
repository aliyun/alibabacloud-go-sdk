// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFolderResponseBody
	GetRequestId() *string
	SetResult(v *DeleteFolderResponseBodyResult) *DeleteFolderResponseBody
	GetResult() *DeleteFolderResponseBodyResult
}

type DeleteFolderResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Response<Map<String, String>>
	Result *DeleteFolderResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFolderResponseBody) GetResult() *DeleteFolderResponseBodyResult {
	return s.Result
}

func (s *DeleteFolderResponseBody) SetRequestId(v string) *DeleteFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFolderResponseBody) SetResult(v *DeleteFolderResponseBodyResult) *DeleteFolderResponseBody {
	s.Result = v
	return s
}

func (s *DeleteFolderResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteFolderResponseBodyResult struct {
	// The request ID.
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	Result map[string]*string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteFolderResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponseBodyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFolderResponseBodyResult) GetResult() map[string]*string {
	return s.Result
}

func (s *DeleteFolderResponseBodyResult) SetRequestId(v string) *DeleteFolderResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DeleteFolderResponseBodyResult) SetResult(v map[string]*string) *DeleteFolderResponseBodyResult {
	s.Result = v
	return s
}

func (s *DeleteFolderResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
