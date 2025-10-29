// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePublicNetworkResponseBody
	GetRequestId() *string
	SetResult(v *UpdatePublicNetworkResponseBodyResult) *UpdatePublicNetworkResponseBody
	GetResult() *UpdatePublicNetworkResponseBodyResult
}

type UpdatePublicNetworkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2A88ECA1-D827-4581-AD39-05149586****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result *UpdatePublicNetworkResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdatePublicNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublicNetworkResponseBody) GetResult() *UpdatePublicNetworkResponseBodyResult {
	return s.Result
}

func (s *UpdatePublicNetworkResponseBody) SetRequestId(v string) *UpdatePublicNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicNetworkResponseBody) SetResult(v *UpdatePublicNetworkResponseBodyResult) *UpdatePublicNetworkResponseBody {
	s.Result = v
	return s
}

func (s *UpdatePublicNetworkResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePublicNetworkResponseBodyResult struct {
	// The status of the public network access switch.
	//
	// example:
	//
	// false
	EnablePublic *bool `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
}

func (s UpdatePublicNetworkResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkResponseBodyResult) GetEnablePublic() *bool {
	return s.EnablePublic
}

func (s *UpdatePublicNetworkResponseBodyResult) SetEnablePublic(v bool) *UpdatePublicNetworkResponseBodyResult {
	s.EnablePublic = &v
	return s
}

func (s *UpdatePublicNetworkResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
