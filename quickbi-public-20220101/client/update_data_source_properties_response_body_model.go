// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourcePropertiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataSourcePropertiesResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateDataSourcePropertiesResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateDataSourcePropertiesResponseBody
	GetSuccess() *bool
}

type UpdateDataSourcePropertiesResponseBody struct {
	// example:
	//
	// D787E1*****DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataSourcePropertiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourcePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourcePropertiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSourcePropertiesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateDataSourcePropertiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataSourcePropertiesResponseBody) SetRequestId(v string) *UpdateDataSourcePropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourcePropertiesResponseBody) SetResult(v bool) *UpdateDataSourcePropertiesResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateDataSourcePropertiesResponseBody) SetSuccess(v bool) *UpdateDataSourcePropertiesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataSourcePropertiesResponseBody) Validate() error {
	return dara.Validate(s)
}
