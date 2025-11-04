// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIngestEndpoints(v []*UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) *UpdateLivePackageChannelCredentialsResponseBody
	GetIngestEndpoints() []*UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints
	SetRequestId(v string) *UpdateLivePackageChannelCredentialsResponseBody
	GetRequestId() *string
}

type UpdateLivePackageChannelCredentialsResponseBody struct {
	// The information about the ingest endpoint.
	IngestEndpoints []*UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints `json:"IngestEndpoints,omitempty" xml:"IngestEndpoints,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 771A1414-27BF-53E6-AB73-EFCB*****ACF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLivePackageChannelCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelCredentialsResponseBody) GetIngestEndpoints() []*UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints {
	return s.IngestEndpoints
}

func (s *UpdateLivePackageChannelCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLivePackageChannelCredentialsResponseBody) SetIngestEndpoints(v []*UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) *UpdateLivePackageChannelCredentialsResponseBody {
	s.IngestEndpoints = v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponseBody) SetRequestId(v string) *UpdateLivePackageChannelCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponseBody) Validate() error {
	if s.IngestEndpoints != nil {
		for _, item := range s.IngestEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints struct {
	// The ingest endpoint ID. `input1` indicates primary and `input2` indicates secondary.
	//
	// example:
	//
	// input1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The password.
	//
	// example:
	//
	// examplePassword123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ingest endpoint URL.
	//
	// example:
	//
	// rtmp://example.com/live/input1
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The username.
	//
	// example:
	//
	// user1
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) GetId() *string {
	return s.Id
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) GetPassword() *string {
	return s.Password
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) GetUrl() *string {
	return s.Url
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) GetUsername() *string {
	return s.Username
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) SetId(v string) *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints {
	s.Id = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) SetPassword(v string) *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints {
	s.Password = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) SetUrl(v string) *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints {
	s.Url = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) SetUsername(v string) *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints {
	s.Username = &v
	return s
}

func (s *UpdateLivePackageChannelCredentialsResponseBodyIngestEndpoints) Validate() error {
	return dara.Validate(s)
}
