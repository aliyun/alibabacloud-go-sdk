// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppCallbackSecretKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackSecretKey(v string) *DescribeAppCallbackSecretKeyResponseBody
	GetCallbackSecretKey() *string
	SetRequestId(v string) *DescribeAppCallbackSecretKeyResponseBody
	GetRequestId() *string
}

type DescribeAppCallbackSecretKeyResponseBody struct {
	// example:
	//
	// a656b296a30xxxxxxxxxx1cd4
	CallbackSecretKey *string `json:"CallbackSecretKey,omitempty" xml:"CallbackSecretKey,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppCallbackSecretKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppCallbackSecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppCallbackSecretKeyResponseBody) GetCallbackSecretKey() *string {
	return s.CallbackSecretKey
}

func (s *DescribeAppCallbackSecretKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppCallbackSecretKeyResponseBody) SetCallbackSecretKey(v string) *DescribeAppCallbackSecretKeyResponseBody {
	s.CallbackSecretKey = &v
	return s
}

func (s *DescribeAppCallbackSecretKeyResponseBody) SetRequestId(v string) *DescribeAppCallbackSecretKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppCallbackSecretKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
