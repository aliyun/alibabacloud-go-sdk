// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutSecretValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutSecretValueResponseBody
	GetRequestId() *string
	SetSecretName(v string) *PutSecretValueResponseBody
	GetSecretName() *string
	SetVersionId(v string) *PutSecretValueResponseBody
	GetVersionId() *string
	SetVersionStages(v *PutSecretValueResponseBodyVersionStages) *PutSecretValueResponseBody
	GetVersionStages() *PutSecretValueResponseBodyVersionStages
}

type PutSecretValueResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// f94ec9d3-2d10-4922-9a5c-5dcd5ebcb5e8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The new version of the secret value.
	//
	// example:
	//
	// 00000000000000000000000000000000203
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that are used to mark the new version.
	VersionStages *PutSecretValueResponseBodyVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
}

func (s PutSecretValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutSecretValueResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *PutSecretValueResponseBody) GetVersionId() *string {
	return s.VersionId
}

func (s *PutSecretValueResponseBody) GetVersionStages() *PutSecretValueResponseBodyVersionStages {
	return s.VersionStages
}

func (s *PutSecretValueResponseBody) SetRequestId(v string) *PutSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutSecretValueResponseBody) SetSecretName(v string) *PutSecretValueResponseBody {
	s.SecretName = &v
	return s
}

func (s *PutSecretValueResponseBody) SetVersionId(v string) *PutSecretValueResponseBody {
	s.VersionId = &v
	return s
}

func (s *PutSecretValueResponseBody) SetVersionStages(v *PutSecretValueResponseBodyVersionStages) *PutSecretValueResponseBody {
	s.VersionStages = v
	return s
}

func (s *PutSecretValueResponseBody) Validate() error {
	return dara.Validate(s)
}

type PutSecretValueResponseBodyVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" type:"Repeated"`
}

func (s PutSecretValueResponseBodyVersionStages) String() string {
	return dara.Prettify(s)
}

func (s PutSecretValueResponseBodyVersionStages) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponseBodyVersionStages) GetVersionStage() []*string {
	return s.VersionStage
}

func (s *PutSecretValueResponseBodyVersionStages) SetVersionStage(v []*string) *PutSecretValueResponseBodyVersionStages {
	s.VersionStage = v
	return s
}

func (s *PutSecretValueResponseBodyVersionStages) Validate() error {
	return dara.Validate(s)
}
