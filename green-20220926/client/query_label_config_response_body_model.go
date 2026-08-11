// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLabelConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContentModeration(v []interface{}) *QueryLabelConfigResponseBody
	GetContentModeration() []interface{}
	SetRequestId(v string) *QueryLabelConfigResponseBody
	GetRequestId() *string
}

type QueryLabelConfigResponseBody struct {
	// The content moderation configuration.
	ContentModeration []interface{} `json:"ContentModeration,omitempty" xml:"ContentModeration,omitempty" type:"Repeated"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryLabelConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLabelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLabelConfigResponseBody) GetContentModeration() []interface{} {
	return s.ContentModeration
}

func (s *QueryLabelConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLabelConfigResponseBody) SetContentModeration(v []interface{}) *QueryLabelConfigResponseBody {
	s.ContentModeration = v
	return s
}

func (s *QueryLabelConfigResponseBody) SetRequestId(v string) *QueryLabelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLabelConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
