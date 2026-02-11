// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAndSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeServiceAndSceneResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *DescribeServiceAndSceneResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DescribeServiceAndSceneResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeServiceAndSceneResponseBodyResultObject) *DescribeServiceAndSceneResponseBody
	GetResultObject() *DescribeServiceAndSceneResponseBodyResultObject
	SetSuccess(v bool) *DescribeServiceAndSceneResponseBody
	GetSuccess() *bool
}

type DescribeServiceAndSceneResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result.
	ResultObject *DescribeServiceAndSceneResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeServiceAndSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAndSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceAndSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeServiceAndSceneResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeServiceAndSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceAndSceneResponseBody) GetResultObject() *DescribeServiceAndSceneResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeServiceAndSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeServiceAndSceneResponseBody) SetCode(v int64) *DescribeServiceAndSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeServiceAndSceneResponseBody) SetHttpStatusCode(v int64) *DescribeServiceAndSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeServiceAndSceneResponseBody) SetRequestId(v string) *DescribeServiceAndSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceAndSceneResponseBody) SetResultObject(v *DescribeServiceAndSceneResponseBodyResultObject) *DescribeServiceAndSceneResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeServiceAndSceneResponseBody) SetSuccess(v bool) *DescribeServiceAndSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeServiceAndSceneResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceAndSceneResponseBodyResultObject struct {
	// Model corresponding scene.
	//
	// example:
	//
	// scene_A
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Model corresponding service.
	//
	// example:
	//
	// service_A
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s DescribeServiceAndSceneResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAndSceneResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeServiceAndSceneResponseBodyResultObject) GetScene() *string {
	return s.Scene
}

func (s *DescribeServiceAndSceneResponseBodyResultObject) GetService() *string {
	return s.Service
}

func (s *DescribeServiceAndSceneResponseBodyResultObject) SetScene(v string) *DescribeServiceAndSceneResponseBodyResultObject {
	s.Scene = &v
	return s
}

func (s *DescribeServiceAndSceneResponseBodyResultObject) SetService(v string) *DescribeServiceAndSceneResponseBodyResultObject {
	s.Service = &v
	return s
}

func (s *DescribeServiceAndSceneResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
