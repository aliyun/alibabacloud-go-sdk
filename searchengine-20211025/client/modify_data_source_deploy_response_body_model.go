// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceDeployResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDataSourceDeployResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyDataSourceDeployResponseBody
	GetResult() map[string]interface{}
}

type ModifyDataSourceDeployResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyDataSourceDeployResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataSourceDeployResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyDataSourceDeployResponseBody) SetRequestId(v string) *ModifyDataSourceDeployResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataSourceDeployResponseBody) SetResult(v map[string]interface{}) *ModifyDataSourceDeployResponseBody {
	s.Result = v
	return s
}

func (s *ModifyDataSourceDeployResponseBody) Validate() error {
	return dara.Validate(s)
}
