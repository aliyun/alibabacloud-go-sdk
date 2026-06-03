// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QuerySceneListResponseBody
	GetRequestId() *string
	SetCode(v string) *QuerySceneListResponseBody
	GetCode() *string
	SetData(v string) *QuerySceneListResponseBody
	GetData() *string
}

type QuerySceneListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QuerySceneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySceneListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySceneListResponseBody) GetData() *string {
	return s.Data
}

func (s *QuerySceneListResponseBody) SetRequestId(v string) *QuerySceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySceneListResponseBody) SetCode(v string) *QuerySceneListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySceneListResponseBody) SetData(v string) *QuerySceneListResponseBody {
	s.Data = &v
	return s
}

func (s *QuerySceneListResponseBody) Validate() error {
	return dara.Validate(s)
}
