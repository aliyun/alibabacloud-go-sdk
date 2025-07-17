// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *TagResourcesResponseBody
	GetData() *string
	SetRequestId(v string) *TagResourcesResponseBody
	GetRequestId() *string
}

type TagResourcesResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 05919CF2-B8A1-588D-B3DB-89B3********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) GetData() *string {
	return s.Data
}

func (s *TagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesResponseBody) SetData(v string) *TagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
