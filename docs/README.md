# :wave: Welcome to Coder's documentation

Find user guides and tutorials, get started with Coder, and more.

<blockquote class="warning">
  <p>
  If you are a Coder v1 customer, view <a href="https://coder.com/docs/coder">the docs</a> or <a href="https://coder.com/docs/coder/latest/guides/v2-faq">the sunset plans</a>.
  </p>
</blockquote>

## Documentation Resources

### :busts_in_silhouette: Developers

Learn how to get started and streamline your coding processes with Coder:

- [**Installation**](https://coder.com/docs/install): Options to install Coder locally, on a VM, or on Kubernetes.
- [**Platforms**](https://coder.com/docs/platforms): Cloud- and platform-specific guides to install and configure a Coder workspace.
- [**Templates**](https://coder.com/docs/templates): Tutorials, guides, and demos for creating and using templates, the backbone of workspaces. Coder workspaces are represented with Terraform, but no Terraform knowledge is required to get started. We have a database of pre-made templates built into the product.
- [**Workspaces**](https://coder.com/docs/workspaces): Instructions for creating, filtering, starting/stopping, and updating workspaces.
- [**IDEs**](https://coder.com/docs/ides): Details on supported Coder IDEs, including desktop and web IDEs, JetBrains Gateway and JetBrains Fleet, Emacs, and remote desktops.
<p align="center">
  <img src="./images/ide-icons.svg" height=72>
</p>

- [**Networking**](https://coder.com/docs/networking): Networking basics and port forwarding, plus advanced networking descriptions of STUN and NAT.
- [**Dotfiles**](https://coder.com/docs/dotfiles): Workspace personalization through dotfiles.
- [**Secrets**](https://coder.com/docs/secrets): Secret management through local configuration and SSH keys. Coder workspaces don't stop at compute. You can add storage buckets, secrets, sidecars
and whatever else Terraform lets you dream up.

For even more resources to help you make the most of Coder, see our [guides](https://coder.com/docs/guides).


## :cloud: Administrators

Coder administrators can control _every_ aspect of their Coder deployment. Learn how to manage your organization's Coder workspaces and resources:

- [**Architecture**](https://coder.com/docs/architecture/architecture): Deployment options, challenges, and risks.
- [**Administration**](https://coder.com/docs/admin): Tutorials and instructions for administering Coder, from creating users and groups to automating, upgrading, and scaling your Coder deployment.
<p align="center">
  <img src="./images/providers-compute.png">
</p>

See our [Templates](https://coder.com/docs/templates) and [Networking](https://coder.com/docs/networking) docs for more details on configuring and securing your organization's deployment.


## Get Started with Coder

Ready to skip the docs and start using Coder now? Use this quickstart to try Coder:

<p align="center">
  <img src="./images/coder-quickstart.png">
</p>

## About Coder

Coder is an open-source platform that lets you create and manage developer workspaces on your public or private cloud infrastructure. It builds on top of common development interfaces (SSH) and infrastructure tools (Terraform) to make **provisioning** and **accessing** remote workspaces simple and efficient.

<p align="center">
  <img src="./images/hero-image.png">
</p>





### Why remote development

Migrating from local developer machines to workspaces hosted by cloud services
is an [increasingly common solution for
developers](https://blog.alexellis.io/the-internet-is-my-computer/) and
[organizations
alike](https://slack.engineering/development-environments-at-slack). There are
several benefits, including:

- **Increased speed:** Server-grade compute speeds up operations in software
  development, such as IDE loading, code compilation and building, and the
  running of large workloads (such as those for monolith or microservice
  applications)

- **Easier environment management:** Tools such as Terraform, nix, Docker,
  devcontainers, and so on make developer onboarding and the troubleshooting of
  development environments easier

- **Increase security:** Centralize source code and other data onto private
  servers or cloud services instead of local developer machines

- **Improved compatibility:** Remote workspaces share infrastructure
  configuration with other development, staging, and production environments,
  reducing configuration drift

- **Improved accessibility:** Devices such as lightweight notebooks,
  Chromebooks, and iPads can connect to remote workspaces via browser-based IDEs
  or remote IDE extensions

### Why Coder

The key difference between Coder OSS and other remote IDE platforms is the added
layer of infrastructure control. This additional layer allows admins to:

- Support ARM, Windows, Linux, and macOS workspaces
- Modify pod/container specs (e.g., adding disks, managing network policies,
  setting/updating environment variables)
- Use VM/dedicated workspaces, developing with Kernel features (no container
  knowledge required)
- Enable persistent workspaces, which are like local machines, but faster and
  hosted by a cloud service

Coder includes [production-ready templates](https://github.com/coder/coder/tree/c6b1daabc5a7aa67bfbb6c89966d728919ba7f80/examples/templates) for use with AWS EC2,
Azure, Google Cloud, Kubernetes, and more.

### What Coder is _not_

- Coder is not an infrastructure as code (IaC) platform. Terraform is the first
  IaC _provisioner_ in Coder, allowing Coder admins to define Terraform
  resources as Coder workspaces.

- Coder is not a DevOps/CI platform. Coder workspaces can follow best practices
  for cloud service-based workloads, but Coder is not responsible for how you
  define or deploy the software you write.

- Coder is not an online IDE. Instead, Coder supports common editors, such as VS
  Code, Vim, and JetBrains, over HTTPS or SSH.

- Coder is not a collaboration platform. You can use Git and dedicated IDE
  extensions for pull requests, code reviews, and pair programming.

- Coder is not a SaaS/fully-managed offering. You must host
  Coder on a cloud service (AWS, Azure, GCP) or your private data center.

## Up next

- Learn about [Templates](./templates/index.md)
- [Install Coder](./install/index.md#install-coder)
