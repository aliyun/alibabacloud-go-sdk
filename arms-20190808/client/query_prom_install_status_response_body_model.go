// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPromInstallStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryPromInstallStatusResponseBodyData) *QueryPromInstallStatusResponseBody
	GetData() *QueryPromInstallStatusResponseBodyData
	SetRequestId(v string) *QueryPromInstallStatusResponseBody
	GetRequestId() *string
}

type QueryPromInstallStatusResponseBody struct {
	// The returned struct.
	Data *QueryPromInstallStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 27E653FA-5958-45BE-8AA9-14D884DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPromInstallStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPromInstallStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPromInstallStatusResponseBody) GetData() *QueryPromInstallStatusResponseBodyData {
	return s.Data
}

func (s *QueryPromInstallStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPromInstallStatusResponseBody) SetData(v *QueryPromInstallStatusResponseBodyData) *QueryPromInstallStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryPromInstallStatusResponseBody) SetRequestId(v string) *QueryPromInstallStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPromInstallStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryPromInstallStatusResponseBodyData struct {
	// Indicates whether the call was successful. Valid values:
	//
	// true: The call was successful. false: The call fails.
	//
	// example:
	//
	// true
	IsControllerInstalled *bool `json:"isControllerInstalled,omitempty" xml:"isControllerInstalled,omitempty"`
}

func (s QueryPromInstallStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPromInstallStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPromInstallStatusResponseBodyData) GetIsControllerInstalled() *bool {
	return s.IsControllerInstalled
}

func (s *QueryPromInstallStatusResponseBodyData) SetIsControllerInstalled(v bool) *QueryPromInstallStatusResponseBodyData {
	s.IsControllerInstalled = &v
	return s
}

func (s *QueryPromInstallStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
