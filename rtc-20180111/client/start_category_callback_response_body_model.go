// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCategoryCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartCategoryCallbackResponseBody
	GetRequestId() *string
}

type StartCategoryCallbackResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartCategoryCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCategoryCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *StartCategoryCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCategoryCallbackResponseBody) SetRequestId(v string) *StartCategoryCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCategoryCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
