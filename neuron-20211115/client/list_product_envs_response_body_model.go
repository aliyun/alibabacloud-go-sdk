// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductEnvsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnvList(v []*string) *ListProductEnvsResponseBody
	GetEnvList() []*string
	SetRequestId(v string) *ListProductEnvsResponseBody
	GetRequestId() *string
}

type ListProductEnvsResponseBody struct {
	EnvList []*string `json:"envList,omitempty" xml:"envList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProductEnvsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductEnvsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductEnvsResponseBody) GetEnvList() []*string {
	return s.EnvList
}

func (s *ListProductEnvsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductEnvsResponseBody) SetEnvList(v []*string) *ListProductEnvsResponseBody {
	s.EnvList = v
	return s
}

func (s *ListProductEnvsResponseBody) SetRequestId(v string) *ListProductEnvsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductEnvsResponseBody) Validate() error {
	return dara.Validate(s)
}
