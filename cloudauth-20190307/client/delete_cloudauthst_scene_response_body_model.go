// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudauthstSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudauthstSceneResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteCloudauthstSceneResponseBody
	GetResultObject() *bool
}

type DeleteCloudauthstSceneResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// 969434DF-926B-4997-9881-4DE94E39F805
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s DeleteCloudauthstSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudauthstSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudauthstSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudauthstSceneResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteCloudauthstSceneResponseBody) SetRequestId(v string) *DeleteCloudauthstSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudauthstSceneResponseBody) SetResultObject(v bool) *DeleteCloudauthstSceneResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteCloudauthstSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
