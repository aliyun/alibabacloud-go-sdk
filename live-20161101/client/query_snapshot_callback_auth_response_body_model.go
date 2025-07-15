// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySnapshotCallbackAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackAuthKey(v string) *QuerySnapshotCallbackAuthResponseBody
	GetCallbackAuthKey() *string
	SetCallbackReqAuth(v string) *QuerySnapshotCallbackAuthResponseBody
	GetCallbackReqAuth() *string
	SetDomainName(v string) *QuerySnapshotCallbackAuthResponseBody
	GetDomainName() *string
	SetRequestId(v string) *QuerySnapshotCallbackAuthResponseBody
	GetRequestId() *string
}

type QuerySnapshotCallbackAuthResponseBody struct {
	// The callback authentication key.
	//
	// example:
	//
	// yourkey
	CallbackAuthKey *string `json:"CallbackAuthKey,omitempty" xml:"CallbackAuthKey,omitempty"`
	// Indicates whether callback authentication is enabled. Valid values:
	//
	// 	- **yes**: Callback authentication is enabled.
	//
	// 	- **no**: Callback authentication is disabled.
	//
	// example:
	//
	// yes
	CallbackReqAuth *string `json:"CallbackReqAuth,omitempty" xml:"CallbackReqAuth,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySnapshotCallbackAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotCallbackAuthResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySnapshotCallbackAuthResponseBody) GetCallbackAuthKey() *string {
	return s.CallbackAuthKey
}

func (s *QuerySnapshotCallbackAuthResponseBody) GetCallbackReqAuth() *string {
	return s.CallbackReqAuth
}

func (s *QuerySnapshotCallbackAuthResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *QuerySnapshotCallbackAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySnapshotCallbackAuthResponseBody) SetCallbackAuthKey(v string) *QuerySnapshotCallbackAuthResponseBody {
	s.CallbackAuthKey = &v
	return s
}

func (s *QuerySnapshotCallbackAuthResponseBody) SetCallbackReqAuth(v string) *QuerySnapshotCallbackAuthResponseBody {
	s.CallbackReqAuth = &v
	return s
}

func (s *QuerySnapshotCallbackAuthResponseBody) SetDomainName(v string) *QuerySnapshotCallbackAuthResponseBody {
	s.DomainName = &v
	return s
}

func (s *QuerySnapshotCallbackAuthResponseBody) SetRequestId(v string) *QuerySnapshotCallbackAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySnapshotCallbackAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
