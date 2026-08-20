// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSessionNetworkRuleTransform interface {
	dara.Model
	String() string
	GoString() string
	SetHeaderValueReplacements(v []*SessionNetworkHeaderValueReplacement) *SessionNetworkRuleTransform
	GetHeaderValueReplacements() []*SessionNetworkHeaderValueReplacement
	SetHeaders(v map[string]*string) *SessionNetworkRuleTransform
	GetHeaders() map[string]*string
}

type SessionNetworkRuleTransform struct {
	// The list of rules for replacing placeholders in HTTP header values before the request is forwarded to the matched host.
	HeaderValueReplacements []*SessionNetworkHeaderValueReplacement `json:"headerValueReplacements" xml:"headerValueReplacements" type:"Repeated"`
	// The HTTP headers injected or overwritten before the request is forwarded to the matched host. Header values are returned in plaintext in GetSession and ListSessions.
	Headers map[string]*string `json:"headers" xml:"headers"`
}

func (s SessionNetworkRuleTransform) String() string {
	return dara.Prettify(s)
}

func (s SessionNetworkRuleTransform) GoString() string {
	return s.String()
}

func (s *SessionNetworkRuleTransform) GetHeaderValueReplacements() []*SessionNetworkHeaderValueReplacement {
	return s.HeaderValueReplacements
}

func (s *SessionNetworkRuleTransform) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SessionNetworkRuleTransform) SetHeaderValueReplacements(v []*SessionNetworkHeaderValueReplacement) *SessionNetworkRuleTransform {
	s.HeaderValueReplacements = v
	return s
}

func (s *SessionNetworkRuleTransform) SetHeaders(v map[string]*string) *SessionNetworkRuleTransform {
	s.Headers = v
	return s
}

func (s *SessionNetworkRuleTransform) Validate() error {
	if s.HeaderValueReplacements != nil {
		for _, item := range s.HeaderValueReplacements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
