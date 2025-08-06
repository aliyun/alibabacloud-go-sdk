// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistPolicyNames(v []*string) *DeleteAppPolicyResponseBody
	GetNonExistPolicyNames() []*string
	SetRequestId(v string) *DeleteAppPolicyResponseBody
	GetRequestId() *string
}

type DeleteAppPolicyResponseBody struct {
	NonExistPolicyNames []*string `json:"NonExistPolicyNames,omitempty" xml:"NonExistPolicyNames,omitempty" type:"Repeated"`
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppPolicyResponseBody) GetNonExistPolicyNames() []*string {
	return s.NonExistPolicyNames
}

func (s *DeleteAppPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppPolicyResponseBody) SetNonExistPolicyNames(v []*string) *DeleteAppPolicyResponseBody {
	s.NonExistPolicyNames = v
	return s
}

func (s *DeleteAppPolicyResponseBody) SetRequestId(v string) *DeleteAppPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
