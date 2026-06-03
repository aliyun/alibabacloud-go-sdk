// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSceneResponseBody
	GetRequestId() *string
	SetCode(v string) *DeleteSceneResponseBody
	GetCode() *string
	SetData(v string) *DeleteSceneResponseBody
	GetData() *string
}

type DeleteSceneResponseBody struct {
	// example:
	//
	// 95E6F720-6786-53BD-9982-C9A636CEA627
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {\\"code\\":\\"200\\",\\"RequestId\\":\\"B946877A-8BDC-55AD-BE7E-85B75B2527E5\\",\\"data\\":true}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSceneResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteSceneResponseBody) SetRequestId(v string) *DeleteSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSceneResponseBody) SetCode(v string) *DeleteSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSceneResponseBody) SetData(v string) *DeleteSceneResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
