// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicWhiteIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePublicWhiteIpsResponseBody
	GetRequestId() *string
	SetResult(v *UpdatePublicWhiteIpsResponseBodyResult) *UpdatePublicWhiteIpsResponseBody
	GetResult() *UpdatePublicWhiteIpsResponseBodyResult
}

type UpdatePublicWhiteIpsResponseBody struct {
	// example:
	//
	// C82758DD-282F-4D48-934F-92170A33****
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdatePublicWhiteIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdatePublicWhiteIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublicWhiteIpsResponseBody) GetResult() *UpdatePublicWhiteIpsResponseBodyResult {
	return s.Result
}

func (s *UpdatePublicWhiteIpsResponseBody) SetRequestId(v string) *UpdatePublicWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicWhiteIpsResponseBody) SetResult(v *UpdatePublicWhiteIpsResponseBodyResult) *UpdatePublicWhiteIpsResponseBody {
	s.Result = v
	return s
}

func (s *UpdatePublicWhiteIpsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePublicWhiteIpsResponseBodyResult struct {
	PublicIpWhitelist []*string `json:"publicIpWhitelist,omitempty" xml:"publicIpWhitelist,omitempty" type:"Repeated"`
}

func (s UpdatePublicWhiteIpsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicWhiteIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePublicWhiteIpsResponseBodyResult) GetPublicIpWhitelist() []*string {
	return s.PublicIpWhitelist
}

func (s *UpdatePublicWhiteIpsResponseBodyResult) SetPublicIpWhitelist(v []*string) *UpdatePublicWhiteIpsResponseBodyResult {
	s.PublicIpWhitelist = v
	return s
}

func (s *UpdatePublicWhiteIpsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
