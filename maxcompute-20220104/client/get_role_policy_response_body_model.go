// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRolePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetRolePolicyResponseBody
	GetData() *string
	SetRequestId(v string) *GetRolePolicyResponseBody
	GetRequestId() *string
}

type GetRolePolicyResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// {
	//
	//       "Statement": [
	//
	//             {
	//
	//                   "Action": [
	//
	//                         "odps:*"
	//
	//                   ],
	//
	//                   "Effect": "Allow",
	//
	//                   "Resource": [
	//
	//                         "acs:odps:*:projects/{projectname}/authorization/packages"
	//
	//                   ]
	//
	//             }
	//
	//       ],
	//
	//       "Version": "1"
	//
	// }
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc1eeed16675342848904412e08dd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRolePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRolePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetRolePolicyResponseBody) GetData() *string {
	return s.Data
}

func (s *GetRolePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRolePolicyResponseBody) SetData(v string) *GetRolePolicyResponseBody {
	s.Data = &v
	return s
}

func (s *GetRolePolicyResponseBody) SetRequestId(v string) *GetRolePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRolePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
